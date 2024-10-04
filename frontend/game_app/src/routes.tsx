import { createBrowserRouter } from "react-router-dom";

import HomePage from "./pages/Home";
import AccessPage from "./pages/Access";
import sendMagicLinkAction from "./actions/sendMagicLink";
import AccessConfirmPage from "./pages/AccessConfirm";
import ScorePage from "./pages/ScorePage";


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
                path: "score-page",
                element: <ScorePage />,
            }
            
        ],
    },

]);

export default router;
