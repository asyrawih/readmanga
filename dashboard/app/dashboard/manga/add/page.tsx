import { AddFormManga } from "@/components/form/addmanga";
import { Card, CardHeader, CardContent } from "@/components/ui/card";

export default function AddManga() {
  return (
    <div className="flex justify-center w-full">
      <Card className="w-full mt-2">
        <CardHeader>
          Add New Manga
        </CardHeader>
        <CardContent>
          <AddFormManga />
        </CardContent>
      </Card>
    </div>
  )
}
