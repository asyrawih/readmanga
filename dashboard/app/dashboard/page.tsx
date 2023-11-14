import { cache } from "react"

type Response = {
  error_code: string
  message: string
  data: Manga[]
}

type Manga = {
  id: number
  title: string
  status: string
  author: string
  type: string
}

export const revalidate = 3600

const getMangas = cache(async () => {
  const result = await fetch(`${process.env.BACKEND_URL}/manga`)
  return result.json() as unknown as Response
})

export default async function Dasboard() {
  const { data } = await getMangas()
  return (
    <>
      {data.map((manga) => (
        <div key={manga.id}>
          {manga.title}
        </div>
      ))}
    </>
  )
} 
