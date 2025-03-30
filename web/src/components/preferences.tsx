import { FC } from "react";
import { useForm } from "react-hook-form"
import { useMediaQuery } from "@uidotdev/usehooks";

import { useAppDispatch, useAppSelector } from "@/store";
import { closeMenu, openMenu } from "@/slices/menu_slice";
import { fetchData, StationFilters, updateFilters } from "@/slices/station_slice";

import { Slider } from "./ui/slider";
import { Button } from "./ui/button";
import { ToggleGroup, ToggleGroupItem } from "./ui/toggle-group";
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from "./ui/drawer"
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "./ui/dialog";

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
  const dispatch = useAppDispatch();
  const brands = useAppSelector(s => s.brands.value);
  const filters = useAppSelector(s => s.stations.filters);

  const { setValue, reset, watch, handleSubmit } = useForm({ defaultValues: filters })
  const filterValues = watch()  

  const handleCancelClick = () => { dispatch(closeMenu('preferences')); reset(); }
  const onData = (data: StationFilters) => {
    dispatch(updateFilters(data));
    dispatch(fetchData());
    dispatch(closeMenu('preferences'))
  }

  return (
    <form className="space-y-10 pt-5">

      <div>
        <h4 className="text-lg">Distance (miles)</h4>
        <div className="flex gap-3 mt-3 mx-2">
          <Slider value={[filterValues.radius]} onValueChange={(vals) => setValue('radius', vals[0])} min={1} max={20} />
          {filterValues.radius}
        </div>
      </div>

      <div>
        <h4 className="text-lg">Brands</h4>
        <ToggleGroup type='multiple' value={filterValues.brands} onValueChange={(vals) => setValue('brands', vals)}>
          {brands.map(b => (
            <ToggleGroupItem key={b} className="m-1 rounded" value={b}>{b}</ToggleGroupItem>
          ))}
        </ToggleGroup>
      </div>

      <div className="text-lg">
        <h4>Fuel Type</h4>
        <ToggleGroup type='multiple' value={filterValues.fuel_types} onValueChange={(vals) => setValue('fuel_types', vals)}>
            <ToggleGroupItem className="m-1 rounded" value='e10'>Petrol (E10)</ToggleGroupItem>
            <ToggleGroupItem className="m-1 rounded" value='e5'>Super (E5)</ToggleGroupItem>
            <ToggleGroupItem className="m-1 rounded" value='b7'>Diesel (B7)</ToggleGroupItem>
            <ToggleGroupItem className="m-1 rounded" value='sdv'>SDV</ToggleGroupItem>
        </ToggleGroup>
      </div>

      <div className="flex gap-2 justify-center">
        <Button type="button" variant='destructive-ghost' onClick={handleCancelClick}>Cancel</Button>
        <Button type="button" variant="primary" onClick={handleSubmit(onData)}>Update</Button>
      </div>

    </form>
  )
}
