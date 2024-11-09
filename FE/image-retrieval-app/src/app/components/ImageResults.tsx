import { Grid, Card, CardMedia, Pagination, Box } from "@mui/material";
import { useState } from "react";

interface ImageResultsProps {
  images: string[];
  limit: number;
  totalImages: number;
  onChangePage: (offset: number) => void;
}

const ImageResults: React.FC<ImageResultsProps> = ({
  images,
  limit,
  totalImages,
  onChangePage,
}) => {
  const [page, setPage] = useState(1);
  const pageCount = Math.ceil(totalImages / limit);

  const handlePageChange = (
    event: React.ChangeEvent<unknown>,
    newPage: number
  ) => {
    onChangePage((newPage - 1) * limit);
    setPage(newPage);
  };

  return (
    <Box>
      <Grid container spacing={2}>
        {(images || []).map((image, index) => (
          <Grid item xs={3} key={index}>
            <Card>
              <CardMedia
                component="img"
                height="140"
                image={`http://localhost:8009/images/${image}`}
                alt={`Image ${index}`}
              />
            </Card>
          </Grid>
        ))}
      </Grid>
      <Box mt={2} display="flex" justifyContent="center">
        <Pagination
          count={pageCount}
          page={page}
          onChange={handlePageChange}
          color="primary"
        />
      </Box>
    </Box>
  );
};

export default ImageResults;
