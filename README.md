### Cloud Run - Deploy movies

This simple setup enables you to serve videos using Cloud Run.

The image that is built is very simple, it consists of two endpoints:
- `/` serves the index.html that references the video.
- `/video` serves the video file in chunks smaller than 32MB to mitigate the max request/response limits on Cloud Run.
    
#### Choose your video file

You can for instance use [Sintel](https://durian.blender.org/download/). Download it, (I chose the 2048-surround version) 
and rename it to `movie.mp4` in the root of the project.
    
#### Build the image    

`gcloud build -t gcr.io/${PROJECT_ID}/movies/sintel .`

#### Push the image to your Google Cloud Repository

`gcloud push gcr.io/${PROJECT_ID}/movies/sintel`

#### Run it on Cloud Run

`gcloud beta run deploy sintel-movie --allow-unauthenticated --image=gcr.io/${PROJECT_ID}/movies/sintel --region=us-central1`

When this deployment is successful, you will be presented with an URL; browse to it and enjoy the video!


#### Deploy it automatically

Register a Google Cloud Build trigger using the provided `cloudbuild.yaml` as starting point.

