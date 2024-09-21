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
                        go version
                        GOOS=linux GOARCH=amd64 go build -o smc
                        zip -r smc-linux-amd64 smc
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
                    // Use o diretório do workspace do Jenkins
                    dir("${env.WORKSPACE}") {  
                        sh '''
                        go version
                        GOOS=windows GOARCH=amd64 go build -o smc.exe
                        zip -r smc-win-amd64 smc.exe
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
