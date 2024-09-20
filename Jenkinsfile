pipeline {
    agent { label 'localhost' }
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
            when {
                anyOf {
                    branch 'develop'
                }
            }
            steps {
                script {
                    dir("${env.WORKSPACE}/Download") {
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
                }
            }
            steps {
                script {
                    dir("${env.WORKSPACE}/Download") { // Usando o diretório correto
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
            echo "Builds completed successfully."
        }
        failure {
            echo "Build failed. Please check the logs."
        }
    }
}
