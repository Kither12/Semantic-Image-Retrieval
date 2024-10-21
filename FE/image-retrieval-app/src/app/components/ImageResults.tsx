import { Grid, Card, CardMedia } from "@mui/material";

interface ImageResult {
  url: string;
}

interface ImageResultsProps {
  images: ImageResult[];
}

const ImageResults: React.FC<ImageResultsProps> = ({ images }) => {
  return (
    <Grid container spacing={2}>
      {images.map((image, index) => (
        <Grid item xs={3} key={index}>
          <Card>
            <CardMedia
              component="img"
              height="140"
              image={image.url}
              alt={`Result image ${index + 1}`}
            />
          </Card>
        </Grid>
      ))}
    </Grid>
  );
};

export default ImageResults;
