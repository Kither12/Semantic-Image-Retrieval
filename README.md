# Semantic Image Retrieval System
# Overview 
This project is a semantic image retrieval system designed to find images based on their content using text query rather than metadata. It supports uploading the image and then listing the images based on the relevance of their content with the text query.
## Demo
https://github.com/user-attachments/assets/8263020a-f117-459c-9f11-d6e70fe0e736
## Key Components
- **Frontend**: Built with Next.js and Material-UI (MUI).
- **Vector databse**: I choose [qdrant](https://github.com/qdrant/qdrant) because it's written in Rust and I love Rust.
- **Embedding model**: Using [clip](https://github.com/openai/CLIP) for fast prototype.
- Then everything are connected using GRPC and the main API gateway are written on Golang.
  ![Screenshot of a comment on a GitHub issue showing an image, added in the Markdown, of an Octocat smiling and raising a tentacle.](/arch.png)



