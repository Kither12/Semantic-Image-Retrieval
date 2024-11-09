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
    return false;
  }
  return true;
}

export async function searchImages({
  prompt,
  limit,
  offset,
}: {
  prompt: string;
  limit: string;
  offset: string;
}) {
  const queryParams = new URLSearchParams({
    prompt,
    limit,
    offset,
  });

  const response = await fetch(
    `http://localhost:8000/api/v1/images?${queryParams.toString()}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!response.ok) {
    throw new Error("Failed to fetch images by text");
  }

  const data = await response.json();
  return {
    path: data.images,
    total: data.total,
  };
}
