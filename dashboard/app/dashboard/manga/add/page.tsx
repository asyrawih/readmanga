import { DropZone } from "@/components/dropImage/image";
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
          <DropZone />
          <AddFormManga />
        </CardContent>
      </Card>
    </div>
  )
}
