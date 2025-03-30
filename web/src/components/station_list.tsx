import { FC, useState } from "react";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useAppDispatch, useAppSelector } from "@/store";

import { Station } from "@/models/station";
import { closeMenu, openMenu } from "@/slices/menu_slice";

import { Button } from "./ui/button";
import StationItem from "./station_item";
import { Sheet, SheetClose, SheetContent } from "./ui/sheet";
import { ToggleGroup, ToggleGroupItem } from "./ui/toggle-group";
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from "./ui/drawer"

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faClose } from "@fortawesome/free-solid-svg-icons";

const snapPoints = ['170px', '600px', 1];

const StationList: FC = () => {
  
  const dispatch = useAppDispatch();
  const [sortBy, setSortBy] = useState<string>('distance');
  const [snap, setSnap] = useState<number|string|null>(snapPoints[0]);

  const isOpen = useAppSelector(s => s.menus.stations);
  const isDesktop = useMediaQuery("(min-width: 768px)")  

  const open = () => dispatch(openMenu('preferences'))
  const close = () => dispatch(closeMenu('stations'))

  const stations = useAppSelector(s => s.stations.value)
  const filteredStations = stations
    .filter(s => !!s[sortBy as keyof Station])
    .sort((a,b) => ((a[sortBy as keyof Station] as number) - (b[sortBy as keyof Station] as number)) || (a.distance - b.distance)) // sort by specified key, if equal, sort by dist
   
  if (isDesktop) {
    return (
      <Sheet open={isOpen} modal={false}>
        <SheetContent side='left' className='z-[10000]' hideClose>
          <SheetClose className="absolute right-4 top-6">
            <Button variant='ghost' size='icon' className="text-2xl size-14 border rounded-full" onClick={close}>
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
          {filteredStations.length == 0 && (
            <div className="p-4 h-full grid">
              <div className="place-self-center space-y-4 text-center">
                <p>Looks like there's nothing here :&#40;</p>
                <Button variant='ghost' onClick={open}>Try changing your preferences</Button>
              </div>
            </div>
          )}
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
    <Drawer snapPoints={snapPoints} activeSnapPoint={snap} setActiveSnapPoint={setSnap} open modal={false} dismissible={false}>
      <DrawerContent className="z-[10000]">
        <DrawerHeader className="text-center text-base">
          <DrawerTitle>{stations.length} stations nearby</DrawerTitle>
        </DrawerHeader>
        <div className="[&>div]:border-b-4 [&>div]:p-4 overflow-y-auto">
          {stations.map(s => (
            <StationItem key={s.site_id} station={s} />
          ))}
        </div>
      </DrawerContent> 
    </Drawer>
  );
};

export default StationList;