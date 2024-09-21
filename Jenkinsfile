pipeline {
    agent { label 'localhost' }
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    stages { 
        stage('Clone') {
            when {
                anyOf {
                    branch 'develop'
                }
            }
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/develop']],
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }
        stage('Build for Linux') {
           agent {
                label getAgentLabel()
            }
            when {
                anyOf {
                    branch 'develop'
                }
            }
            steps {
                script {
                    // Use o diretório do workspace do Jenkins
                    dir("${env.WORKSPACE}") {  
                        sh '''
                        go version  # Verifica se o Go está disponível
                        GOOS=linux GOARCH=amd64 go build -o nome-do-app-linux
                        '''
                    }
                }
            }
        }
    }
    post {
        success {
            echo "Builds completed successfully."
        }
        failure {
            echo "Build failed. Please check the logs."
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