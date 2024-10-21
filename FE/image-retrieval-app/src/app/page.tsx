// pages/index.tsx
"use client";
import { Container, Typography } from "@mui/material";
import SearchBar from "@/app/components/SearchBar";
import { uploadImage, searchImages } from "./actions/uploadImageAction";

export default function Home() {
  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        Image Retrieval System
      </Typography>
      <SearchBar onSearch={searchImages} onUpload={uploadImage} />
    </Container>
  );
}
