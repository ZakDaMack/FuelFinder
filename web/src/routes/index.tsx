import HomePage from "@/pages";
import AuthLayout from "@/pages/auth";
import StationPage from "@/pages/station";
import LoginPage from "@/pages/auth/login_form";
import SignupPage from "@/pages/auth/signup_form";

const routes = [
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: 'auth',
    element: <AuthLayout />,
    children: [
      {
        path: 'register',
        element: <SignupPage />,
      },
      {
        path: 'login',
        element: <LoginPage />,
      }
      // {
      //   path: 'register',
      //   element: <RegisterPage />,
      // }
    ]
  },
  {
    path: "/stations/:id",
    element: <StationPage />,
  },
];

export { routes };