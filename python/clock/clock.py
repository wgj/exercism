class Clock(object):
    def __init__(self, hour, minute):
        self.hour = hour
        self.minute = minute

    def __str__(self):
        return '%.2d:%.2d' % (self.hour, self.minute)

    # Print time
    def add(self, minutes):
        self.minute = self.minute + minutes
