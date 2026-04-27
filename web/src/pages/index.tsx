import Map from "@/components/map";
import Init from "@/components/init";
import IntroDialog from "@/components/intro_dialog";
import Preferences from "@/components/preferences";

const HomePage = () => {
  return (
    <main>
      <IntroDialog />
      <Preferences />
      <Init />
      <Map />
    </main>
  );
}

export default HomePage;