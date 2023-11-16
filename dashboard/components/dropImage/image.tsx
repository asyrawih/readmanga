'use client'

import { useMutation } from "react-query"
import { useDropzone } from "react-dropzone"
import { CSSProperties, useEffect, useState } from "react";


type ImageFile = File & {
  preview: string
}

export const DropZoneComponent = () => {
  const [files, setFiles] = useState<Array<ImageFile>>([]);
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

  const thumbs = files.map(file => (
    <div key={file.name} className="mt-3 mb-3 flex space-x-1">
      <div className="w-[400px] m-2">
        <img
          src={file.preview}
          // Revoke data uri after image is loaded
          onLoad={() => { URL.revokeObjectURL(file.preview) }}
        />
      </div>
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

