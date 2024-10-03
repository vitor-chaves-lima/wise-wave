import { createBrowserRouter } from "react-router-dom";
import HomePage from "./pages/Home";
import AccessPage from "./pages/Access";
import sendMagicLinkAction from "./actions/sendMagicLink";
import AccessConfirmPage from "./pages/AccessConfirm";
import LastGamePage from "./pages/LastGame";
import GameExamplePage from "./pages/GameExample";

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
                path: "last-game",
                element: <LastGamePage />,
            },
            {
                path: "game",
                element: <GameExamplePage />,
            }
        ],
    },
]);

export default router;
