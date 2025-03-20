import { FC, useState } from "react";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useAppDispatch, useAppSelector } from "@/store";
import { closeMenu, openMenu } from "@/slices/menu_slice";

import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import {
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  DrawerTitle,
  DrawerTrigger,
} from "@/components/ui/drawer"
import { Drawer } from 'vaul'
import StationItem from "./station_item";
import { ToggleGroup, ToggleGroupItem } from "./ui/toggle-group";
import { Station } from "@/models/station";
import { Button } from "./ui/button";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faClose } from "@fortawesome/free-solid-svg-icons";

const snapPoints = ['80px', '450px', 0.88];

const StationList: FC = () => {
  
  const dispatch = useAppDispatch();
  const [sortBy, setSortBy] = useState<string>('distance');
  const [snap, setSnap] = useState<number|string|null>(snapPoints[0]);

  const isOpen = useAppSelector(s => s.menus.stations);
  const isDesktop = useMediaQuery("(min-width: 768px)")  

  const open = () => dispatch(openMenu('stations'))
  const close = () => dispatch(closeMenu('stations'))
  const setOpen = () => isOpen ? close() : open();

  const stations = useAppSelector(s => s.stations.value)
  const filteredStations = stations
    .filter(s => !!s[sortBy as keyof Station])
    .sort((a,b) => ((a[sortBy as keyof Station] as number) - (b[sortBy as keyof Station] as number)) || (a.distance - b.distance)) // sort by specified key, if equal, sort by dist
   
  if (isDesktop) {
    return (
      <Sheet open={isOpen} onOpenChange={setOpen}>
        <SheetContent side='left' className='z-[10000]' hideClose>
          <SheetClose className="absolute right-4 top-8">
            <Button variant='ghost' size='icon' className="text-xl size-12 border rounded-full">
              <FontAwesomeIcon icon={faClose} />
            </Button>
          </SheetClose>
          <div className="pt-24 px-4">
            <p>Sort by:</p>
            <ToggleGroup type='single' value={sortBy} onValueChange={setSortBy}>
                <ToggleGroupItem className="m-1 ml-0 rounded" value='distance'>Distance</ToggleGroupItem>
                <ToggleGroupItem className="m-1 rounded" value='e10'>Petrol</ToggleGroupItem>
                <ToggleGroupItem className="m-1 rounded" value='e5'>Super</ToggleGroupItem>
                <ToggleGroupItem className="m-1 rounded" value='b7'>Diesel</ToggleGroupItem>
            </ToggleGroup>
          </div>
          <div className="border-t [&>div]:border-b-4 [&>div]:p-4 overflow-y-auto">
            {filteredStations.map(s => (
              <StationItem key={s.site_id} station={s} />
            ))}
          </div>
        </SheetContent>
      </Sheet>
    )
  }
   
  return (
    <Drawer.Root snapPoints={snapPoints} activeSnapPoint={snap} setActiveSnapPoint={setSnap} open dismissible={false}>
      {/* <DrawerTrigger>ttest</DrawerTrigger> */}
      <Drawer.Content className='z-[10000] relative'>
        {/* <DrawerHeader className="text-left">
          <DrawerTitle>Stations ({stations.length})</DrawerTitle>
        </DrawerHeader> */}
        <div className="[&>div]:border-b-4 [&>div]:p-4 overflow-y-auto">
          {stations.map(s => (
            <StationItem key={s.site_id} station={s} />
          ))}
        </div>
      </Drawer.Content >
    </Drawer.Root>
  );
};

export default StationList;