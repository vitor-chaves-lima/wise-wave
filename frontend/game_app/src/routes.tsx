import { createBrowserRouter } from "react-router-dom";
import HomePage from "./pages/Home";
import AccessPage from "./pages/Access";
import sendMagicLinkAction from "./actions/sendMagicLink";
import AccessConfirmPage from "./pages/AccessConfirm";
import LastGamePage from "./pages/LastGame";
import GameStartPage from "./pages/GameStart";
import GameExamplePage from "./pages/GameExample";
import MagicLinkValidatePage from "./pages/MagicLinkValidate";
import validateMagicLink from "./loaders/validateMagicLink.ts";
import MagicLinkValidateErrorBoundary from "./components/MagicLinkValidateErrorBoundary.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        children: [
            {
				index: true,
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
				path: "magic-link/validate",
				element: <MagicLinkValidatePage />,
				loader: validateMagicLink,
				errorElement: <MagicLinkValidateErrorBoundary />
			},
            {
                path: "last-game",
                element: <LastGamePage />,
            },
            {
                path: 'game-start',
                element: <GameStartPage/>
            },
            {
                path: "game",
                element: <GameExamplePage />,
            }
        ],
    },
]);

export default router;
