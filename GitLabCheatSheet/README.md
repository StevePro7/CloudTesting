# GitLab Cheat Sheet
TODO - move this somewhere else

base64
https://cloudifydevops.com/secured-gitlab-pipeline-with-base64

  script:
    - echo "$SECRET_KEY" | base64 -d
    - export DECODED_SECRET_KEY="$SECRET_KEY"
	
	
https://gitlab.com/StevePro7/hello-world-python/-/blob/main/.gitlab-ci.yml	