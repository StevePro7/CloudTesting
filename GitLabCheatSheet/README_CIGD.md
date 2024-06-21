ReadMe CICD
15-Jun-2024

https://medium.com/@rubyabdullah14/how-to-auto-build-and-push-to-dockerhub-using-gitlab-ci-22408e0ca51b
https://stackoverflow.com/questions/51196435/gitlab-ci-docker-command-not-found


Publish a Docker image to Docker Hub using GitLab
https://dev.to/mattdark/publish-a-docker-image-to-docker-hub-using-gitlab-2b9o


CI_REGISTRY             docker.io
CI_REGISTRY_IMAGE       
CI_REGISTRY_PASSWORD    SP4
CI_REGISTRY_USER        steveproxna


docker login -u "$CI_REGISTRY_USER" -p "$CI_REGISTRY_PASSWORD" $CI_REGISTRY

THIS WORKED
docker login -u "steveproxna" -p "$CI_REGISTRY_PASSWORD" docker.io

maybe three is an issue with either 
$CI_REGISTRY_USER
$CI_REGISTRY


Previos

docker build --pull --rm -f "Dockerfile" -t flask-api:latest "."
docker tag flask-api:latest steveproxna/flask-api:1.0
docker push steveproxna/flask-api:1.0


gitlab deploy helm to cluster example

Deploy an application with GitLab CI/CD
https://www.dbi-services.com/blog/deploy-an-application-with-gitlab-ci-cd/

- sed -i "s/#TAG#/${CI_COMMIT_SHORT_SHA}/g" ./charts/values.yaml


      ${HELM_ADDITIONAL_ARGS}
      helm upgrade
      --install
      ${HELM_RELEASE_NAME}
      ${HELM_CHART_PATH}
      --kubeconfig ${KUBECONFIG}
      --namespace ${DEPLOYMENT_NAMESPACE}
      --description "${DEPLOYMENT_DESCRIPTION}"
      ${HELM_ADDITIONAL_ARGS}



build job: chosen stage does not exist; available stages are .pre, deploy, .post