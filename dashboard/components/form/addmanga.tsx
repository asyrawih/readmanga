'use client'

import { useForm } from "react-hook-form"
import * as z from "zod"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "../ui/form"
import { Input } from "../ui/input"
import { Separator } from "../ui/separator"
import { Textarea } from "../ui/textarea"
import { Button } from "../ui/button"


// {
//   "author": "string",
//   "release_date": "string",
//   "sinopsis": "string",
//   "status": "string",
//   "title": "string",
//   "total_chapter": 0,
//   "type": "string"
// }
const formSchema = z.object({
  author: z.string({ required_error: "author name required" }).min(3).max(200),
  release_date: z.string({ required_error: "Release Date required" }),
  sinopsis: z.string({ required_error: "sinopsis required" }),
  status: z.string({ required_error: "status required" }),
  title: z.string({ required_error: "title required" }),
  total_chapter: z.number({ required_error: "total_chapter required" }),
  type: z.string({ required_error: "type of manga required" })
})

type FormUI = {
  name: keyof z.infer<typeof formSchema>
  label: string
  type: string
}

const buildForm: Array<FormUI> = [
  {
    name: 'title',
    label: "Title",
    type: "text"
  },
  {
    name: "author",
    label: "Author",
    type: "text"
  },
  {
    name: "release_date",
    label: "Release Date",
    type: "text"
  },
  {
    name: "type",
    label: "Type",
    type: "text"
  },
  {
    name: "total_chapter",
    label: "Total Chapter",
    type: "text"
  },
  {
    name: "status",
    label: "Status",
    type: "text"
  },
  {
    name: "sinopsis",
    label: "Sinopsis",
    type: "textarea"
  },
]

export const AddFormManga = () => {
  const form = useForm<z.infer<typeof formSchema>>({
    defaultValues: {
      ...(buildForm.reduce((acc, item) => {
        acc[item.name] = "";
        return acc;
      }, {} as { [key: string]: string })),
    }
  })

  const handleSubmit = (val: z.infer<typeof formSchema>) => {
    console.log(val)
  }

  return (
    <>
      <Form {...form}>
        <form className="w-full" onSubmit={form.handleSubmit(handleSubmit)}>
          {buildForm.map((item) => (
            <FormField
              key={item.name}
              control={form.control}
              name={item.name}
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    {item.type == "textarea" ? (<Textarea {...field} />) : (<Input className="my-3" placeholder={item.label} autoComplete="off" {...field} />)}
                  </FormControl>
                </FormItem>
              )}
            />
          )
          )}
          <Button type="submit">Save</Button>
        </form>
      </Form>
    </>
  )
}
