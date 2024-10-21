// app/layout.tsx
import "./globals.css";
import { ReactNode } from "react";
import { CssBaseline } from "@mui/material";

export const metadata = {
  title: "Image Retrieval System",
  description: "A simple image retrieval system using Next.js and MUI",
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body>
        <CssBaseline />
        {children}
      </body>
    </html>
  );
}
