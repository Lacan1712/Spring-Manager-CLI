pipeline {
    agent any
    stages {
        stage('Clone') {
            when {
                anyOf {
                    branch 'develop'
                    branch 'feature/*'
                }
            }
            steps {
                checkout([$class: 'GitSCM', 
                    branches: [[name: '*/develop']], // Você pode alterar isso se quiser que a branch específica mude.
                    userRemoteConfigs: [[url: 'https://github.com/Lacan1712/Spring-Manager-CLI.git']]
                ])
            }
        }
        stage('Build for Linux') {
            when {
                anyOf {
                    branch 'develop'
                    branch 'feature/*'
                }
            }
            steps {
                script {
                    dir('~/Download') { // Altere para o diretório correto onde o Go está
                        sh '''
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
                    branch 'feature/*'
                }
            }
            steps {
                script {
                    dir('~/Download') { // Altere para o diretório correto onde o Go está
                        sh '''
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
