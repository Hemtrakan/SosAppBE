#!groovy

def ENVIRONMENT = 'dev'


def PROJECT_ENV = 'dev'
def PROJECT_ACCOUNT = 'accounts'
def PROJECT_EMERGENCY = 'emergency'
def PROJECT_HOTLINE = 'hotline'
def PROJECT_MESSENGER = 'messenger'
def DIR_PROJECT = '/home/chirapon_job/SosAppBE'


pipeline {
	agent any

	environment {
        image1 = "annabells/sosapp-accounts-service"
        image2 = "annabells/sosapp-emergency-service"
        image3 = "annabells/sosapp-hotline-service"
        image4 = "annabells/sosapp-messenger-service"
        registry = "docker.io"
    }


	stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Print Environment') {
            steps {
                sh('ls -al')
                sh('printenv')
            }
        }

		stage('Build docker image') {
            steps {
                script {
                    docker.withRegistry('', 'dockerhub') {
                        dir("./${PROJECT_ACCOUNT}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image1}:${BUILD_NUMBER} ."
                        }

                        dir("./${PROJECT_EMERGENCY}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image2}:${BUILD_NUMBER} ."
                        }

                        dir("./${PROJECT_HOTLINE}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image3}:${BUILD_NUMBER} ."
                        }

                        dir("./${PROJECT_MESSENGER}"){
                            sh "sed -i 's/{ENV}/${PROJECT_ENV}/g' Dockerfile"
                            sh "docker build -t ${env.image4}:${BUILD_NUMBER} ."
                        }

                        sh("docker push ${env.image1}:latest")
                        sh("docker push ${env.image2}:latest")
                        sh("docker push ${env.image3}:latest")
                        sh("docker push ${env.image4}:latest")

                        sh('docker logout')
                    }
                }
            }
        }

		stage('Deployment'){
            steps {
                dir("./${DIR_PROJECT}"){
                    sh "docker-compose up -d"
                }
            }
        }
	}
}