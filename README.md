##Image-processing-service
This service provides you a graphql API for an image resizing 

Link - https://image-processing-5hpdgtorsa-lz.a.run.app/query

##Schema is [here](https://github.com/PatrickKvartsaniy/image-processing-service/blob/master/graph/schema.graphqls)
I highly recommend you to use [Insomnia](https://insomnia.rest/]) for graphql requests 

##If you prefer curl
curl https://image-processing-5hpdgtorsa-lz.a.run.app/query \
  -F operations={"query":"mutation($file: Upload!){upload(image:$file,parameters:{width:100,height:100}) {id  path type size ts variety{path width height}}}","variables":{}}' \
  -F map='{"f":["variables.file"]}' \
  -F f={path_to_your_image_here}


##Deployment
Service has been deployed to the Google Cloud Platform using [Cloud Run](https://cloud.google.com/run)

##CI/CD
by GitHub and  [Cloud Build](https://cloud.google.com/cloud-build) 

##Database
Service uses [Atlas](https://cloud.google.com/cloud-build) MongoDB as the main database.

##Storage
Images stored in [Cloud Storage](https://cloud.google.com/storage)
