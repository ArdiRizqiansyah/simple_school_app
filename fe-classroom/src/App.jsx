import { useHydration } from "./hooks/useHydration";
import LoadingPage from "./pages/loadingPage";
import Routes from "./routes"

function App() {
  const { loading } = useHydration();

  if (loading) {
    return (
      <LoadingPage/>
    );
  }

  return (
    <Routes />
  )
}

export default App
