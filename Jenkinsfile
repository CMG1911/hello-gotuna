pipeline {
    agent any
    options {
      timestamps ()
      ansiColor('xterm')
    }  
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

