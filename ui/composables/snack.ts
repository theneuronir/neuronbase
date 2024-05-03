type Queue = Item[];

interface Item {
  title: string;
  value?: boolean;
  message?: string;
  color?: string;
  timeout?: number;
}

interface NotifyOpts extends Queue {}

export const useSnack = () => {
  const snack = useState("snack", () => <Queue>[]);

  const notify = (opts: NotifyOpts) => {
    for (const opt of opts) {
      snack.value.push({ ...opt, value: true });
      setTimeout(() => {
        snack.value.pop();
      }, opt.timeout || 5000);
    }
  };

  return { snack, notify };
};
