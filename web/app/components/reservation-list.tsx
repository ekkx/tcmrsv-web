import { ReservationListItem } from "./reservation-list-item";

export function ReservationList() {
  return (
    <div className="grid gap-6 px-4">
      <div className="grid gap-2">
        <h4 className="ml-2 text-sm text-foreground-400">6月</h4>
        <div className="flex flex-col gap-2">
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
        </div>
      </div>
      <div className="grid gap-2">
        <h4 className="ml-2 text-sm text-foreground-400">7月</h4>
        <div className="flex flex-col gap-2">
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
        </div>
      </div>
      <div className="grid gap-2">
        <h4 className="ml-2 text-sm text-foreground-400">8月</h4>
        <div className="flex flex-col gap-2">
          <ReservationListItem />
          <ReservationListItem />
        </div>
      </div>
      <div className="grid gap-2">
        <h4 className="ml-2 text-sm text-foreground-400">9月</h4>
        <div className="flex flex-col gap-2">
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
          <ReservationListItem />
        </div>
      </div>
    </div>
  );
}
