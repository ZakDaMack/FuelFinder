import { FC } from "react";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useAppDispatch, useAppSelector } from "@/store";
import { closeMenu, openMenu } from "@/slices/menu_slice";

import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import {
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerTitle,
} from "@/components/ui/drawer"
import StationItem from "./station_item";

const StationList: FC = () => {
  
  const dispatch = useAppDispatch();

  const isOpen = useAppSelector(s => s.menus.stations);
  const isDesktop = useMediaQuery("(min-width: 768px)")  

  const open = () => dispatch(openMenu('stations'))
  const close = () => dispatch(closeMenu('stations'))
  const setOpen = () => isOpen ? close() : open();

  const stations = useAppSelector(s => s.stations.value)
   
  if (isDesktop) {
    return (
      <Sheet open={isOpen} onOpenChange={setOpen}>
        <SheetContent side='bottom'>
          <SheetHeader>
            <SheetTitle>Stations ({stations.length})</SheetTitle>
          </SheetHeader>
          <div>
            {stations.map(s => (
              <StationItem key={s.site_id} station={s} />
            ))}
          </div>
        </SheetContent>
      </Sheet>
    )
  }
   
  return (
    <Drawer open={isOpen} onOpenChange={setOpen}>
      <DrawerContent>
        <DrawerHeader className="text-left">
          <DrawerTitle>Stations ({stations.length})</DrawerTitle>
        </DrawerHeader>
        <div>
          {stations.map(s => (
            <StationItem key={s.site_id} station={s} />
          ))}
        </div>
      </DrawerContent>
    </Drawer>
  );
};

export default StationList;