'use client'

import { useMutation } from "react-query"

const uploadThumbnail = async () => {
}


export const DropZone = () => {

  const mutate = useMutation(uploadThumbnail, {
    onSuccess: () => { },
    onError: () => { }
  })

  return (
    <div>Dropzone in here</div>
  )
}

