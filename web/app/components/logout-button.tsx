import { Button } from "@heroui/react";
import { Cookie } from "~/store/cookies";

export function LogoutButton() {
  return (
    <Button
      isIconOnly
      variant="light"
      size="lg"
      className="rounded-full"
      onPress={() => {
        Cookie.destroy();
        window.location.href = "/login";
      }}
      startContent={
        <svg
          className="w-6 h-6 text-foreground-500"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
        >
          <path
            fill="currentColor"
            d="M12 3.25a.75.75 0 0 1 0 1.5a7.25 7.25 0 0 0 0 14.5a.75.75 0 0 1 0 1.5a8.75 8.75 0 1 1 0-17.5"
          />
          <path
            fill="currentColor"
            d="M16.47 9.53a.75.75 0 0 1 1.06-1.06l3 3a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06l1.72-1.72H10a.75.75 0 0 1 0-1.5h8.19z"
          />
        </svg>
      }
    />
  );
}
