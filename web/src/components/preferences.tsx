import { FC } from "react";
import { cn } from "@/lib/utils";
import { useForm } from "react-hook-form"
import { useMediaQuery } from "@uidotdev/usehooks";

import { useAppDispatch, useAppSelector } from "@/store";
import { closeMenu, openMenu } from "@/slices/menu_slice";
import { boundsSelector, fetchData, StationFilters, updateFilters } from "@/slices/station_slice";

import { Button } from "./ui/button";
import BrandLogo from "./brand_logo";
import { GradientFuelIcon } from "./gradient_fuel_icon";
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

const availableFuelTypes = [
  { value: 'e10', label: 'Petrol', colour: 'from-green-400 to-green-700' },
  { value: 'e5', label: 'Super', colour: 'from-blue-400 to-blue-700' },
  { value: 'b7', label: 'Diesel', colour: 'from-yellow-400 to-yellow-700' },
  { value: 'sdv', label: 'Super Diesel', colour: 'from-orange-400 to-orange-700' },
]

const PreferencesForm: FC = () => {
  const dispatch = useAppDispatch();
  const bounds = useAppSelector(boundsSelector);
  const brands = useAppSelector(s => s.brands.value);
  const filters = useAppSelector(s => s.stations.filters);

  const { setValue, reset, watch, handleSubmit } = useForm({ defaultValues: filters })
  const filterValues = watch()  

  const handleCancelClick = () => { dispatch(closeMenu('preferences')); reset(); }

  const onData = (data: StationFilters) => {
    dispatch(updateFilters(data));
    dispatch(fetchData(bounds!));
    dispatch(closeMenu('preferences'))
  }

  return (
    <form className="space-y-10 pt-5">
      {/* <div>
        <h4 className="text-lg">Distance (miles)</h4>
        <div className="flex gap-3 mt-3 mx-2">
          <Slider value={[filterValues.radius]} onValueChange={(vals) => setValue('radius', vals[0])} min={1} max={20} />
          {filterValues.radius}
        </div>
      </div> */}

      <div>
        <div className='flex items-center'>
          <h4 className="text-lg">Brands</h4>
          <span className="text-neutral-600 text-sm pl-2">({filterValues.brands.length == 0 ? 'All' : filterValues.brands.length})</span>
          <div className="grow" />
          <Button variant='ghost' type='button' size='sm' onClick={() => setValue('brands', [])}>Clear</Button>
        </div>
        <ToggleGroup type='multiple' value={filterValues.brands} onValueChange={(vals) => setValue('brands', vals)}>
          {brands.map(b => (
            <ToggleGroupItem key={b} className={cn(
              "size-20 p-px m-1 rounded-[16px] from-primary from-60% to-blue-400 shadow-primary/50",
              "data-[state=on]:bg-linear-to-r data-[state=on]:shadow data-[state=on]:text-black",
            )} value={b}>
              <div className="rounded-[15px] grid place-items-center [&>p]:text-sm! [&>img]:w-4/5 bg-white size-full">
                <BrandLogo brand={b} />
              </div>
            </ToggleGroupItem>
          ))}
        </ToggleGroup>
      </div>

      <div>
        <div className='flex items-center'>
          <h4 className="text-lg">Fuel Type</h4>
          <span className="text-neutral-600 text-sm pl-2">({filterValues.fuel_types.length == 0 ? 'All' : filterValues.fuel_types.length})</span>
          <div className="grow" />
          <Button variant='ghost' type='button' size='sm' onClick={() => setValue('fuel_types', [])}>Clear</Button>
        </div>
        <ToggleGroup type='multiple' value={filterValues.fuel_types} onValueChange={(vals) => setValue('fuel_types', vals)}>
            {availableFuelTypes.map(f => (
              <ToggleGroupItem key={f.value} value={f.value} className={cn(
                "size-20 p-px m-1 rounded-[16px]", f.colour,
                "data-[state=on]:bg-linear-to-r data-[state=on]:shadow",
              )}>
                <div className="rounded-[15px] grid relative bg-white size-full">
                  <GradientFuelIcon className={cn("place-self-center size-6 bg-gradient-to-r", f.colour)} />
                  <p className="text-xs text-black">{f.label}</p>
                </div>
              </ToggleGroupItem>
            ))}
        </ToggleGroup>
      </div>

      <div className="flex gap-2 justify-center">
        <Button type="button" variant='destructive-ghost' onClick={handleCancelClick}>Cancel</Button>
        <Button type="button" variant="primary" onClick={handleSubmit(onData)}>Update</Button>
      </div>

    </form>
  )
}
