import {
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { ReservationForm } from "./reservation-form";

type Props = {
  isConfirmed?: boolean;
  rooms?: any[];
  reservationId: number;
  campusType: string;
  date: string;
  timeRange: string;
  roomName: string;
  roomId?: string;
  reservation?: any;
  userName?: string;
};

export function EditReservationButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  // Check if reservation is within 3 days
  const isWithin3Days = () => {
    // Use the actual date from reservation object
    const dateString = props.reservation?.date;
    
    if (!dateString) {
      return false; // If no date available, allow editing
    }
    
    const reservationDate = new Date(dateString);
    const today = new Date();
    
    // Set both dates to start of day for accurate comparison
    today.setHours(0, 0, 0, 0);
    reservationDate.setHours(0, 0, 0, 0);
    
    const diffTime = reservationDate.getTime() - today.getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    
    return diffDays < 3;
  };

  const cannotEdit = props.isConfirmed || isWithin3Days();

  return (
    <>
      <Button fullWidth onPress={onOpen} size="sm" variant="flat">
        編集
      </Button>
      {cannotEdit ? (
        <Modal
          isOpen={isOpen}
          onOpenChange={onOpenChange}
          placement="center"
          size="xs"
          closeButton={<></>}
        >
          <ModalContent>
            {(onClose) => (
              <>
                <ModalBody className="p-0 gap-0">
                  <div className="grid gap-4 py-6 text-center">
                    <p className="text-xl font-bold">予約を変更できません</p>
                    <p className="text-xs">
                      この予約は確定しているため、変更できません。
                    </p>
                  </div>
                  <div className="flex justify-center gap-6 border-t py-3">
                    <Button
                      className="w-32 font-bold border"
                      variant="bordered"
                      onPress={onClose}
                    >
                      閉じる
                    </Button>
                  </div>
                </ModalBody>
              </>
            )}
          </ModalContent>
        </Modal>
      ) : (
        <Drawer
          isOpen={isOpen}
          onOpenChange={onOpenChange}
          size="xl"
          classNames={{
            closeButton: "top-4 right-4 scale-125 border",
          }}
        >
          <DrawerContent>
            {(onClose) => (
              <>
                <DrawerBody className="gap-8 py-8">
                  <div className="grid gap-1">
                    <h3 className="text-xl font-bold">予約変更</h3>
                    <p className="text-xs text-foreground-400">
                      希望するセクションを変更してください。
                    </p>
                  </div>
                  <ReservationForm 
                    type="update" 
                    rooms={props.rooms || []} 
                    reservationId={props.reservation?.id || props.reservationId}
                    defaultCampus={props.reservation?.campusType === 2 ? "ikebukuro" : "nakameguro"}
                    defaultDate={props.reservation?.date}
                    defaultStartTime={props.reservation ? `${props.reservation.fromHour}:${String(props.reservation.fromMinute).padStart(2, "0")}` : props.timeRange.split(" 〜 ")[0].replace(/^0(\d):/, '$1:')}
                    defaultEndTime={props.reservation ? `${props.reservation.toHour}:${String(props.reservation.toMinute).padStart(2, "0")}` : props.timeRange.split(" 〜 ")[1].replace(/^0(\d):/, '$1:')}
                    defaultRoomName={props.reservation?.roomId || props.roomId}
                    defaultUserName={props.reservation?.bookerName || props.userName}
                    onSuccess={onClose}
                  />
                </DrawerBody>
              </>
            )}
          </DrawerContent>
        </Drawer>
      )}
    </>
  );
}
