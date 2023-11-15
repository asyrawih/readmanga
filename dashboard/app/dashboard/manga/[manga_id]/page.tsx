'use client'
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { BACKEND_URL } from "@/lib/utils";
import { Separator } from "@radix-ui/react-separator";
import Link from "next/link";
import { useQuery } from "react-query";

type Response = {
  error_code: string
  message: string
  data: {
    Manga: Manga
    Chapters: Chapter[]
  }
}
type Chapter = {
  id: number,
  manga_id: number,
  chapter: string
}

type Manga = {
  id: number
  title: string
  status: string
  author: string
  type: string
}

export default function DetailPage({ params }: { params: { manga_id: string } }) {
  const getMangaDetail = async () => {
    const result = await fetch(`${BACKEND_URL}/manga/${params.manga_id}`)
    return result.json()
  }
  const { data } = useQuery<Response, Error>("manga_detail", getMangaDetail)


  if (!data) {
    return
  }

  const chapters = data.data.Chapters?.sort((a, b) => Number(a) - Number(b))

  return (
    <div>
      <Card>
        <CardHeader>
          <CardTitle className="mt-2 inline-flex flex-col">
            <span className="my-2 text-muted font-bold tracking-wide text-lg">Title</span>
            {data.data.Manga.title}
          </CardTitle>
          <CardTitle className="mt-2  inline-flex flex-col">
            <span className="my-2 text-muted font-bold tracking-wide text-lg">Status</span>
            {data.data.Manga.status}
          </CardTitle>
          <CardTitle className="mt-2  inline-flex flex-col">
            <span className="my-4 text-muted font-bold tracking-wide text-xl">Type</span>
            {data.data.Manga.type}
          </CardTitle>
          <CardTitle className="mt-2 inline-flex flex-col">
            <span className="my-2 text-muted font-bold tracking-wide text-xl">Author</span>
            {data.data.Manga.author}
          </CardTitle>
          <Button variant={'secondary'}>Add Chapter</Button>
        </CardHeader>
        <CardContent>
          <CardTitle>Chapter List</CardTitle>
          <Separator orientation="horizontal" className="my-3 border" />
          <ScrollArea className="h-[450px]">
            {chapters?.map(item => (
              <div key={item.id}>
                <div className="border px-2 py-3 cursor-pointer">
                  <Link href={`/dashboard/manga/${item.manga_id}/read/${item.id}`}>
                    chapter {item.chapter}
                  </Link>
                </div>
              </div>
            ))}
          </ScrollArea>
        </CardContent>
      </Card>
    </div>
  )
}
