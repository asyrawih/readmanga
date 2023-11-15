'use client'

import { Button, buttonVariants } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { ToastAction } from "@/components/ui/toast"
import { toast } from "@/components/ui/use-toast"
import { BACKEND_URL } from "@/lib/utils"
import { CardStackPlusIcon, EyeOpenIcon, TableIcon, TrashIcon } from "@radix-ui/react-icons"
import Link from "next/link"
import { useMutation, useQuery, useQueryClient } from "react-query"

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

const deleteManga = async (id: number) => {
  const result = await fetch(`${BACKEND_URL}/manga/${id}`, {
    method: "DELETE",
    headers: {
      'Content-Type': "application/json"
    }
  })
  return result.json()
}

export default function Manga() {
  const getMangas = async () => {
    const result = await fetch(`${BACKEND_URL}/manga`)
    return result.json()
  }
  const query = useQueryClient()
  const { data } = useQuery<Response, Error>("mangas", getMangas)

  const mutate = useMutation(deleteManga, {
    onSuccess: () => {
      query.invalidateQueries({ queryKey: ["mangas"] })
    }
  })

  const sendDeleteRequest = (manga: Manga) => {
    // mutate.mutate(manga.id)
    toast({
      title: `Want Delete ${manga.title} ?`,
      variant: "destructive",
      action: <ToastAction className="outline" onClick={() => mutate.mutate(manga.id)} altText="delete">Delete ?</ToastAction>,
    })
  }

  if (!data) {
    return (
      <>
        Loading ...
      </>
    )
  }

  if (query.isFetching({ queryKey: ["mangas"] })) {
    return (
      <>
        Fetching ....
      </>
    )
  }

  return (
    <Card className="mt-2">
      <CardHeader className="flex justify-between flex-row">
        <CardTitle>Manga List</CardTitle>
        <Button variant={'secondary'}>
          <Link href={'/dashboard/manga/add'}> Add manga </Link>
        </Button>
      </CardHeader>
      <CardContent>
        <Table>
          <TableCaption>Manga List</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px]">id</TableHead>
              <TableHead>Title</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Type</TableHead>
              <TableHead className="text-right">Action</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {data.data?.map(item => (
              <TableRow key={item.id}>
                <TableCell>{item.id} </TableCell>
                <TableCell>{item.title} </TableCell>
                <TableCell>{item.status} </TableCell>
                <TableCell>{item.type} </TableCell>
                <TableCell className="text-right space-x-1">
                  <Link href={`/dashboard/manga/${item.id}`} className={buttonVariants({ variant: 'outline' })}>
                    <EyeOpenIcon className="mx-1" />
                  </Link>
                  <Button disabled={mutate.status == 'loading'} variant={'destructive'} onClick={() => sendDeleteRequest(item)}>
                    <TrashIcon className="mx-1" />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
