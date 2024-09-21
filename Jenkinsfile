pipeline {
    agent none // Define que não haverá um agent padrão
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    stages {
        stage('Determine Agent') {
            agent {
                // Define o agent dinamicamente com base na branch
                label getAgentLabel()
            }
            steps {
                script {
                    echo "Agent selected for branch: ${env.BRANCH_NAME}"
                }
            }
        }
        stage('Clone') {
            agent {
                label getAgentLabel() // Chama a função para obter o label do agent
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
                label getAgentLabel() // Chama a função para obter o label do agent
            }
            steps {
                script {
                    // Comandos de build aqui
                    echo "Building for branch: ${env.BRANCH_NAME}"
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

// Função para determinar o label do agent com base na branch
def getAgentLabel() {
    switch (env.BRANCH_NAME) {
        case 'develop':
            return 'develop-agent' // Define o label para a branch develop
        case { it.startsWith('feature/') }:
            return 'feature-agent' // Define o label para branches de feature
        case { it.startsWith('release/') }:
            return 'release-agent' // Define o label para branches de release
        default:
            return 'default-agent' // Define um label padrão para outras branches
    }
}
