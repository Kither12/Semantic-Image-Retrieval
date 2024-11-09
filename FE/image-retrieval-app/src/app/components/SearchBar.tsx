"use client";
import { useState, ChangeEvent, createRef } from "react";
import {
  TextField,
  Button,
  Grid,
  Typography,
  Stack,
  Alert,
} from "@mui/material";

interface SearchBarProps {
  onSearch: (formData: FormData) => void;
  onUpload: (formData: FormData) => Promise<boolean>;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch, onUpload }) => {
  const [query, setQuery] = useState<string>("");
  const [imageName, setImageName] = useState<string>("");
  const [uploadError, setUploadError] = useState<string | null>(null);
  const ref = createRef<HTMLFormElement>();

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      setImageName(files[0].name);
      setUploadError(null); // Reset error when a new file is selected
    }
  };

  const handleSubmitClient = async (formData: FormData) => {
    const response = await onUpload(formData);
    if (response) {
      // success
      ref.current?.reset();
      setImageName("");
      setUploadError(null);
    } else {
      // fail
      setUploadError("Failed to upload image. Please try again.");
    }
  };

  return (
    <Grid container spacing={2} alignItems="center">
      {/* Search Form */}
      <Grid item xs={12} sm={8}>
        <form action={onSearch}>
          <Stack direction="row" alignItems="center" gap={2}>
            <TextField
              fullWidth
              label="Search by text"
              variant="outlined"
              name="prompt"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
            />
            <Button type="submit" variant="contained" color="primary">
              Search
            </Button>
          </Stack>
        </form>
      </Grid>

      {/* Image Upload Form */}
      <Grid item xs={12} sm={4}>
        <form action={handleSubmitClient} ref={ref}>
          <input
            accept="image/*"
            type="file"
            name="image"
            onChange={handleFileChange}
            style={{ display: "none" }}
            id="upload-image"
          />
          <label htmlFor="upload-image">
            <Button variant="contained" component="span">
              Upload Image
            </Button>
          </label>
          {imageName && (
            <>
              <Typography variant="body2" color="textSecondary">
                Uploaded image: {imageName}
              </Typography>
              <Button
                type="submit"
                variant="contained"
                color="primary"
                style={{ marginTop: "10px" }}
              >
                Submit Image
              </Button>
            </>
          )}
          {uploadError && (
            <Alert severity="error" style={{ marginTop: "10px" }}>
              {uploadError}
            </Alert>
          )}
        </form>
      </Grid>
    </Grid>
  );
};

export default SearchBar;
