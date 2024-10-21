"use server";

export async function uploadImage(formData: FormData) {
  const file = formData.get("image") as File;
  const buffer = Buffer.from(await file.arrayBuffer());
  const image_data = buffer.toString("base64");
  const payload = {
    image: image_data,
    fileName: file.name,
    contentType: file.type,
  };

  const response = await fetch("http://localhost:8000/api/v1/images", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    const errorMessage = await response.text();
    console.error("Image upload failed: " + errorMessage);
    return false;
  }

  return true;
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export async function searchImages(formData: FormData) {
  // const response = await fetch(
  //   `http://localhost:8000/api/v1/images?query=${encodeURIComponent(query)}`,
  //   {
  //     method: "GET",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //   }
  // );
  // if (!response.ok) {
  //   throw new Error("Failed to fetch images by text");
  // }
  // const data = await response.json();
  // return data.images;
}
