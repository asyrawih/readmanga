'use client'

import { Button, buttonVariants } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { BACKEND_URL } from "@/lib/utils"
import { CardStackPlusIcon, TableIcon } from "@radix-ui/react-icons"
import Link from "next/link"
import { useQuery, useQueryClient } from "react-query"

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

export default function Manga() {
  const getMangas = async () => {
    const result = await fetch(`${BACKEND_URL}/manga`)
    return result.json()
  }
  const query = useQueryClient()
  const { data } = useQuery<Response, Error>("mangas", getMangas)

  if (!data) {
    return (
      <>
        No Data For Showing
      </>
    )
  }

  return (
    <Card className="mt-12">
      <CardHeader className="flex justify-between flex-row">
        <CardTitle>Manga List</CardTitle>
        <Button variant={'secondary'}>Add Manga</Button>
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
            {data.data.map(item => (
              <TableRow key={item.id}>
                <TableCell>{item.id} </TableCell>
                <TableCell>{item.title} </TableCell>
                <TableCell>{item.status} </TableCell>
                <TableCell>{item.type} </TableCell>
                <TableCell className="text-right">
                  <Link href={`/dashboard/manga/${item.id}`} className={buttonVariants({ variant: 'ghost' })}>
                    View
                  </Link>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
