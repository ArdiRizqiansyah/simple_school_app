import Navbar from "@Components/navbar";

import PropTypes from "prop-types";
import { Toaster } from "react-hot-toast";

export default function UserLayout({ children }) {
  return (
    <div className="min-h-screen bg-slate-100">
      <Navbar />
      <div className="container max-w-screen-xl px-4 py-5 mx-auto">
        {children}
      </div>

      <Toaster
        position="top-right"
        reverseOrder={false}
      />
    </div>
  );
}

UserLayout.propTypes = {
  children: PropTypes.node.isRequired,
};
