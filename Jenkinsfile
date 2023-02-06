pipeline {
    agent any
    options {
      timestamps()
        stages {
          stage('Build') {
            steps {
             sh 'docker-compose build'
                }
          }
          stage('Up') {
            steps {
             sh 'docker-compose up -d'
            }
          }
        }
      }
    }
}
