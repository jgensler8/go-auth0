#!/usr/bin/env bash

# !!!
# v2.3.0 doesn't contain https://github.com/go-swagger/go-swagger/commit/b49f5203edce35bd92c8f01e410352829c8c6f18#diff-d7872b02ccb665c39db474d558af68fd
# !!!
# echo "INFO : *** Running swagger-codegen in a docker container ***"
# docker run -it \
#   -v $PWD/scripts/swagger.yaml:/swagger.yaml \
#   -v $PWD/generated:/generated \
#   jimschubert/swagger-codegen-cli \
#   generate \
#   -i /swagger.yaml \
#   -l go \
#   -o /generated
# if [ $? -eq 1 ]; then
#   echo "ERROR : Docker run for swagger code generation failed"
# else
#   echo "INFO : SUCCESSFUL code generation "
# fi

mkdir -p generated
SWAGGER=$GOPATH/bin/swagger
$SWAGGER generate client \
  --spec=scripts/swagger.yaml \
  --target=generated
