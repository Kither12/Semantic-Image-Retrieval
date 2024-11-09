// pages/index.tsx
"use client";
import { Box, Container, Stack, Typography } from "@mui/material";
import SearchBar from "@/app/components/SearchBar";
import { uploadImage, searchImages } from "./actions/uploadImageAction";
import ImageResults from "./components/ImageResults";
import { useEffect, useState } from "react";

const limit = 12; //page_limit

export default function Home() {
  const [searchResults, setSearchResults] = useState([]);
  const [prompt, setPrompt] = useState("");
  const [offset, setOffset] = useState(0);
  const [total, setTotal] = useState(0);

  const handleSearch = (formData: FormData) => {
    setPrompt(formData.get("prompt")?.toString() || "");
  };

  const handlerUpload = async (formData: FormData) => {
    const res = await uploadImage(formData);
    if (res === false) {
      return false;
    }
    async function fetchImages() {
      const res = await searchImages({
        prompt,
        limit: limit.toString(),
        offset: offset.toString(),
      });
      setSearchResults(res.path);
      setTotal(res.total);
    }
    await fetchImages();
    return true;
  };

  const onChangePage = (offset: number) => {
    setOffset(offset);
  };

  useEffect(() => {
    async function fetchImages() {
      const res = await searchImages({
        prompt,
        limit: limit.toString(),
        offset: offset.toString(),
      });
      setSearchResults(res.path);
      setTotal(res.total);
    }
    fetchImages();
  }, [prompt, offset]);

  return (
    <Container
      maxWidth={false}
      sx={{
        height: "100vh",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography variant="h4" gutterBottom sx={{ m: 3 }}>
        Image Retrieval System
      </Typography>
      <Stack spacing={2} sx={{ width: "70%" }}>
        <Box
          sx={{
            width: "100%",
            display: "flex",
            justifyContent: "center",
          }}
        >
          <SearchBar onSearch={handleSearch} onUpload={handlerUpload} />
        </Box>
        <ImageResults
          images={searchResults}
          limit={limit}
          onChangePage={onChangePage}
          totalImages={total}
        />
      </Stack>
    </Container>
  );
}
