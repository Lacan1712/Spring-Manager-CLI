pipeline {
    agent none // Define que não haverá um agent padrão
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
        SCRIPTS_PATH = "${env.WORKSPACE}/scripts"
    }
    stages {
        stage('Clone') {
            agent {
                label getAgentLabel()
            }
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: "${env.BRANCH_NAME}"]],
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }
        stage('Build') {
            agent {
                label getAgentLabel()
            }
            steps {
                script {
                    // Define o comando a ser executado baseado na branch
                    def buildCommand = ''
                    // Executa o script de build correspondente com base na branch
                    switch (env.BRANCH_NAME) {
                        case 'main':
                            sh "${SCRIPTS_PATH}/build_main.sh"
                            break
                        case 'develop':
                            sh "${SCRIPTS_PATH}/scripts/builds/build_develop.sh"
                            break
                        case { it.startsWith('feature/') }:
                            sh "${SCRIPTS_PATH}/build_feature.sh"
                            break
                        case { it.startsWith('release/') }:
                            sh "${SCRIPTS_PATH}/build_release.sh"
                            break
                        default:
                            sh "${SCRIPTS_PATH}/build_default.sh"
                            break
                    }

                    echo "Building for branch: ${env.BRANCH_NAME}"
                    sh buildCommand
                    
                }
            }
        }
    }
    post {
        success {
            echo "Build completed successfully for branch ${env.BRANCH_NAME}."
        }
        failure {
            echo "Build failed for branch ${env.BRANCH_NAME}. Please check the logs."
        }
    }
}

def getAgentLabel() {
    switch (env.BRANCH_NAME) {
        case 'develop':
            return 'localhost' // Define o label para a branch develop
        case { it.startsWith('feature/') }:
            return 'feature-agent' // Define o label para branches de feature
        case { it.startsWith('release/') }:
            return 'release-agent' // Define o label para branches de release
        default:
            return 'default-agent' // Define um label padrão para outras branches
    }
}
