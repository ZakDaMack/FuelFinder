import { FC } from "react";
import { cn } from "@/lib/utils";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useAppDispatch, useAppSelector } from "@/store";
import { closeMenu, openMenu } from "@/slices/menu_slice";

import {
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerTitle,
} from "@/components/ui/drawer"
import { 
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle
} from "./ui/dialog";

import { Slider } from "./ui/slider";
import { Button } from "./ui/button";
import { ToggleGroup, ToggleGroupItem } from "./ui/toggle-group";


const Preferences: FC = () => {

  const dispatch = useAppDispatch();

  const isOpen = useAppSelector(s => s.menus.preferences);
  const isDesktop = useMediaQuery("(min-width: 768px)")

  const open = () => dispatch(openMenu('preferences'))
  const close = () => dispatch(closeMenu('preferences'))
  const setOpen = () => isOpen ? close() : open();
   
  if (isDesktop) {
    return (
      <Dialog open={isOpen} onOpenChange={setOpen} >
        <DialogContent className='z-[10000]'>
          <DialogHeader>
            <DialogTitle>Preferences</DialogTitle>
          </DialogHeader>
          <PreferencesForm />
        </DialogContent>
      </Dialog>
    )
  }
   
  return (
    <Drawer open={isOpen} onOpenChange={setOpen}>
      <DrawerContent className='z-[10000] p-4'>
        <DrawerHeader>
          <DrawerTitle>Preferences</DrawerTitle>
        </DrawerHeader>
        <PreferencesForm />
      </DrawerContent>
    </Drawer>
  )
};

export default Preferences;

 
const PreferencesForm: FC = () => {
  const brands = useAppSelector(s => s.brands.value);
  const filters = useAppSelector(s => s.stations.filters);

  return (
    <form className="space-y-5 pt-5">

      <div>
        <h4>Distance (miles)</h4>
        <div className="flex gap-3 mt-3">
          <Slider value={[filters.radius]} min={1} max={20} />
          {filters.radius}
        </div>
      </div>

      <div>
        <h4>Brands</h4>
        <ToggleGroup type='multiple'>
          {brands.map(b => (
            <ToggleGroupItem value={b}>{b}</ToggleGroupItem>
          ))}
        </ToggleGroup>
      </div>

      <div>
        <h4>Fuel Type</h4>
        <ToggleGroup type='multiple'>
            <ToggleGroupItem className="rounded-full" value='e10'>Petrol</ToggleGroupItem>
            <ToggleGroupItem value='e5'>Super (E5)</ToggleGroupItem>
            <ToggleGroupItem value='b7'>Diesel</ToggleGroupItem>
            <ToggleGroupItem value='sdv'>SDV</ToggleGroupItem>
        </ToggleGroup>
      </div>

      <div className="flex gap-4 justify-center">
        <Button variant='destructive'>Cancel</Button>
        <Button>Update</Button>
      </div>

    </form>
  )
}
