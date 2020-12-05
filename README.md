API documentation: 
    response: 
        application/json:
            400: 
                description: invalid param
                schema:
                    type: object
                    properties:
                        error:
                            type: string
                            description: error message
            200:
                description: largest prime smaller than given param
                schema:
                    type: object
                    properties:
                        result:
                            type: integer

Start up: 
    - cd to folder
    - run command "docker build -t go-test:latest . && docker run -p 3000:3000 go-test:latest"

Cloud:
    url: http://development.eba-vzvy4htv.us-west-2.elasticbeanstalk.com//{param}
    provider: aws
    service: elastic beanstalk
Local:
    url: http://localhost:3000/{param}


