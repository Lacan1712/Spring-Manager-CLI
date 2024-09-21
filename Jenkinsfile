pipeline {
    label getAgentLabel() 
    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    stages {
        stage('Clone') {
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: "${env.BRANCH_NAME}"]],
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }
        stage('Build for Linux') {
            when {
                anyOf {
                    branch 'develop'
                    branch 'feature/*' // Adicione outras branches conforme necessário
                }
            }
            steps {
                script {
                    dir("${env.WORKSPACE}") {
                        sh '''
                        go version  # Verifica se o Go está disponível
                        GOOS=linux GOARCH=amd64 go build -o nome-do-app-linux
                        '''
                    }
                }
            }
        }
        stage('Build for Windows') {
            when {
                anyOf {
                    branch 'develop'
                    branch 'release/*' // Adicione outras branches conforme necessário
                }
            }
            steps {
                script {
                    dir("${env.WORKSPACE}") {
                        sh '''
                        go version  # Verifica se o Go está disponível
                        GOOS=windows GOARCH=amd64 go build -o nome-do-app-windows.exe
                        '''
                    }
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