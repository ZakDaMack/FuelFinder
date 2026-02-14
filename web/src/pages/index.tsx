import Map from "@/components/map";
import Init from "@/components/init";
import IntroDialog from "@/components/intro_dialog";
import Preferences from "@/components/preferences";
import StationList from "@/components/station_list";

const HomePage = () => {
  return (
    <main>
      <IntroDialog />
      <Preferences />
      <StationList />
      <Init />
      <Map />
    </main>
  );
}

export default HomePage;