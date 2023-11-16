'use client'

import { useMutation } from "react-query"
import { useDropzone } from "react-dropzone"
import { CSSProperties, useEffect, useState } from "react";
import { useEffect, useState } from "react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { TrashIcon } from "@radix-ui/react-icons";
import { BACKEND_URL } from "@/lib/utils";


interface ImageFile extends File {
  id: number,
  preview: string
}

const uploadBatch = async (images: ImageFile[]) => {
  const formData = new FormData()

  for (let i = 0; i < images.length; i++) {
    formData.append('images', images[i]);
  }

  formData.append("model_id", "12")
  formData.append("model_type", "manga")
  formData.append("manga", "test")
  formData.append("chapter", "12")

  const result = await fetch(`${BACKEND_URL}/media/batch`, {
    method: "POST",
    body: formData,
  })

  console.log(await result.json())

}

export const DropZoneComponent = ({ ...props }) => {
  const [files, setFiles] = useState<Array<ImageFile>>([]);

  const [selected, setSelected] = useState<string | null>()

  const handleRemove = (image: ImageFile, index: number) => {
    const newData = [...files]
    newData.splice(index, 1)
    setFiles(newData)
  }

  const { getRootProps, getInputProps } = useDropzone({
    accept: {
      'image/jpeg': [],
      'image/png': [],
      'image/webp': []
    },
    maxFiles: 2,
    onDrop: acceptedFiles => {
      setFiles(acceptedFiles.map(file => Object.assign(file, {
        preview: URL.createObjectURL(file)
      })))
    }
  });

  const thumbs = files.map((file, index) => (
    <div key={file.name} className="mt-3 mb-3 flex space-x-1 relative">
      <div className="w-[400px] m-2">
        <img
          src={file.preview}
          // Revoke data uri after image is loaded
          onLoad={() => { URL.revokeObjectURL(file.preview) }}
        />
      </div>
      <Button className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2" variant={'default'} onClick={() => handleRemove(file, index)}>
        <TrashIcon />
      </Button>
    </div>
  ));

  useEffect(() => {
    // Make sure to revoke the data uris to avoid memory leaks, will run on unmount
    return () => files.forEach(file => URL.revokeObjectURL(file.preview));
  }, []);

  return (
    <>
      <div className="flex">
        {thumbs}
      </div>

      {files.length == 0 && (
        <section className="container my-2 outline-dashed  outline-cyan-800 p-12 flex justify-center  ">
          <div {...getRootProps({ className: 'dropzone' })}>
            <>
              <Input {...getInputProps()} />
              <p>Drag ur thumbnail in here </p>
            </>
          </div>
        </section>
      )}
    </>
  );
}

