"use client";
import { useState, ChangeEvent, createRef } from "react";
import { TextField, Button, Grid, Typography } from "@mui/material";

interface SearchBarProps {
  onSearch: (formData: FormData) => Promise<void>;
  onUpload: (formData: FormData) => Promise<boolean>;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch, onUpload }) => {
  const [query, setQuery] = useState<string>("");
  const [imageName, setImageName] = useState<string>("");
  const ref = createRef<HTMLFormElement>();

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      setImageName(files[0].name); // Set the image name to display it
    }
  };

  const handleSubmitClient = async (formData: FormData) => {
    const response = await onUpload(formData);
    if (response) {
      // success
      ref.current?.reset();
    } else {
      // fail
    }
  };

  return (
    <Grid container spacing={2} alignItems="center">
      {/* Search Form */}
      <Grid item xs={12} sm={8}>
        <form action={onSearch}>
          <TextField
            fullWidth
            label="Search by text"
            variant="outlined"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            style={{ marginTop: "10px" }}
          >
            Search
          </Button>
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
        </form>
      </Grid>
    </Grid>
  );
};

export default SearchBar;
