'use client'

import { useQuery } from "react-query"


type ParamsProps = {
  params: {
    manga_id: number,
    chapter_id: number
  }
}

type Response = {
  error_code: string
  message: string
  data: Model
}

type Model = {
  chapter: Chapter
  medias: Media[]
}

type Chapter = {
  id: number,
  chapter: string
  content: string
}

type Media = {
  id: number,
  model_type: string
  url: string
}

export default function ChapterDetail({ params }: ParamsProps) {
  const getChapterIdDetail = async () => {
    const result = await fetch(`http://localhost:8000/chapter/${params.chapter_id}`)
    return result.json()
  }

  const { data } = useQuery<Response, Error>("getDetailChapters", getChapterIdDetail)

  if (!data) {
    return
  }

  if (!data.data.medias) {
    return (
      <>
        Belum Ada Chapter
      </>
    )
  }

  const medias = data.data.medias?.sort((a, b) => {
    return a.id - b.id
  })

  return (
    <>
      {medias.map(item => (
        <img key={item.id} src={`http://localhost:9000/manga/${item.url}`} />
      ))}
    </>
  )
}
