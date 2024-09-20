pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                script {
                    // Defina suas branches específicas
                    def branches = ['main', 'develop', 'feature/*']

                    // Verifique se a branch atual está na lista
                    if (branches.any { it == env.BRANCH_NAME || it.endsWith("/*") && env.BRANCH_NAME.startsWith(it.substring(0, it.length() - 2)) }) {
                        echo "Executando script na branch: ${env.BRANCH_NAME}"
                        // Adicione seu script de build aqui
                    } else {
                        echo "Branch ${env.BRANCH_NAME} não está na lista, pulando."
                    }
                }
            }
        }
    }
}
