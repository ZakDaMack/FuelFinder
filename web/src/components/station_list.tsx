import { FC, useMemo, useState } from "react";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useAppDispatch, useAppSelector } from "@/store";

import { Station } from "@/models/station";
import { closeMenu, openMenu } from "@/slices/menu_slice";

import { Button } from "./ui/button";
import StationItem from "./station_item";
import { Sheet, SheetClose, SheetContent } from "./ui/sheet";
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from "./ui/drawer"
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuLabel, DropdownMenuTrigger, DropdownMenuRadioItem, DropdownMenuRadioGroup } from "./ui/dropdown-menu";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowDownShortWide, faClose, faSort } from "@fortawesome/free-solid-svg-icons";

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
  const filteredStations = useMemo(() => {
    return stations
      .filter(s => !!s[sortBy as keyof Station] && s[sortBy as keyof Station] as number > 0)
      .sort((a,b) => ((a[sortBy as keyof Station] as number) - (b[sortBy as keyof Station] as number)) || ((a.distance ?? 0) - (b.distance ?? 0)))
      // sort by specified key, if equal, sort by dist
  }, [stations, sortBy]) 
   
  const handleHover = (enter: boolean, siteId: string) => {
    if (!isDesktop) return;
    const el = document.querySelector(`div[data-site-id='${siteId}']`)
    const classes = ['scale-120', '-translate-y-3', 'shadow-xl', 'shadow-black/30', 'z-[10000]!'];
    enter ? el?.classList.add(...classes) : el?.classList.remove(...classes);
  }

  if (isDesktop) {
    return (
      <Sheet open={isOpen} modal={false}>
        <SheetContent side='left' className='z-[10000]' hideClose>
          <SheetClose className="absolute right-4 top-6">
            <Button variant='ghost' size='icon' className="text-2xl size-14 border rounded-full" onClick={close}>
              <FontAwesomeIcon icon={faClose} />
            </Button>
          </SheetClose>
          <div className="pt-24 px-4 flex justify-between items-center">
            <p className='text-lg font-semibold pl-2'>{filteredStations.length} stations nearby</p>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="outline">
                  <FontAwesomeIcon className="mr-2" icon={faArrowDownShortWide} />
                  <FontAwesomeIcon fontSize={10} icon={faSort} />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="w-40 z-[10000]" align="end">
                <DropdownMenuGroup>
                  <DropdownMenuLabel>Sort By</DropdownMenuLabel>
                   <DropdownMenuRadioGroup value={sortBy} onValueChange={setSortBy}>
                    <DropdownMenuRadioItem value="distance">Distance</DropdownMenuRadioItem>
                    <DropdownMenuRadioItem value="e10">Petrol</DropdownMenuRadioItem>
                    <DropdownMenuRadioItem value="e5">Super</DropdownMenuRadioItem>
                    <DropdownMenuRadioItem value="b7">Diesel</DropdownMenuRadioItem>
                  </DropdownMenuRadioGroup>
                </DropdownMenuGroup>
              </DropdownMenuContent>
            </DropdownMenu>
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
              <StationItem key={s.site_id} station={s} onHover={(enter) => handleHover(enter, s.site_id)} />
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