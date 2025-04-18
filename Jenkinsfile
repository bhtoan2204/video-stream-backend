pipeline {
  agent any

  options {
      skipDefaultCheckout(true)
      disableConcurrentBuilds()
  }

  environment {
    HARBOR_USERNAME = credentials('HARBOR_USERNAME')
    HARBOR_PASSWORD = credentials('HARBOR_PASSWORD')
    HARBOR_HOST = credentials('HARBOR_HOST')
    HARBOR_PROJECT = 'youtube'
    TAG = "latest"
  }



  stages {
    stage('Docker Version') {
      steps {
        script {
          docker.withTool('docker') {
            sh 'docker --version'
          }
        }
      }
    }

    stage('Checkout') {
      steps {
        script {
          def gitVars = checkout scm
          COMMIT_ID = gitVars.GIT_COMMIT.take(7)
        }
      }
    }

    stage('Build Images') {
      steps {
        script {
          docker.withTool('docker') {
            def services = ['gateway', 'user', 'video', 'comment', 'worker']
              for (svc in services) {
                def imageName = "${env.HARBOR_HOST}/${env.HARBOR_PROJECT}/${svc}:${env.TAG}"
                def buildContext = "./apps/${svc}"
                def dockerfile = "${buildContext}/Dockerfile"
                sh "docker build -t ${imageName} -f ${dockerfile} ${buildContext}"
              }
          }
        }
      }
    }

    stage('Login to Harbor') {
      steps {
        script {
          docker.withTool('docker') {
            sh "docker login ${env.HARBOR_HOST} -u ${env.HARBOR_CRED} -p ${env.HARBOR_PASS}"
          }
        }
      }
    }

    stage('Push Images') {
      steps {
        script {
          docker.withTool('docker') {
            def services = ['gateway', 'user', 'video', 'comment', 'worker']
            for (svc in services) {
              def image = "${env.HARBOR_HOST}/${env.HARBOR_PROJECT}/${svc}:${env.TAG}"
              sh "docker push ${image}"
            }
          }
        }
      }
    }

    stage('Clean up') {
      steps {
        script {
          docker.withTool('docker') {
            sh 'docker system prune -af'
          }
        }
      }
    }
  }
}