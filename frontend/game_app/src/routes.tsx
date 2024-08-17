import { createBrowserRouter } from "react-router-dom";

import HomePage from "./pages/Home";
import AccessPage from "./pages/Access";

const router = createBrowserRouter([
    {
        path: "/",
        children: [
            {
                path: "/",
                element: <HomePage />,
            },
            {
                path: "access",
                element: <AccessPage />,
            }
        ],
    },
]);

export default router;
