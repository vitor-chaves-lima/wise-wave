import { createBrowserRouter } from "react-router-dom";

import HomePage from "./pages/Home";
import AccessPage from "./pages/Access";
import sendMagicLinkAction from "./actions/sendMagicLink";
import AccessConfirmPage from "./pages/AccessConfirm";
import AntesUltimaPage from "./pages/AntesUltima";

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
                action: sendMagicLinkAction,
            },
            {
                path: "access-confirm",
                element: <AccessConfirmPage />,
            },
            {
                path: "antes-ultima",
                element: <AntesUltimaPage />,
            }
        ],
    },

]);

export default router;
