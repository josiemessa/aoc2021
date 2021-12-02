using System;
using System.IO;

namespace Aoc21
{
    public class Day2
    {
        public static int Part1(string path)
        {
            int horizontal = 0;
            int depth = 0;
            foreach (var line in File.ReadLines(path))
            {
                var split = line.Split(" ");
                var value = int.Parse(split[1]);
                switch (split[0])
                {
                    case "forward":
                        horizontal += value;
                        break;
                    case "down":
                        depth += value;
                        break;
                    case "up":
                        depth -= value;
                        break;
                }
            }

            return horizontal * depth;
        }

        public static int Part2(string path)
        {
            int horizontal = 0;
            int depth = 0;
            int aim = 0;
            foreach (var line in File.ReadLines(path))
            {
                var split = line.Split(" ");
                var value = int.Parse(split[1]);
                switch (split[0])
                {
                    case "forward":
                        horizontal += value;
                        depth += aim * value;
                        break;
                    case "down":
                        aim += value;
                        break;
                    case "up":
                        aim -= value;
                        break;
                }
            }

            return horizontal * depth;
        }
    }
}